package resources

import (
	"regexp"
	"strconv"

	"github.com/cockroachdb/errors"
	"go.mondoo.io/mondoo/lumi"
	"go.mondoo.io/mondoo/lumi/resources/tlsshake"
)

var reTarget = regexp.MustCompile("([^/:]+?)(:\\d+)?$")

func (s *lumiTls) init(args *lumi.Args) (*lumi.Args, Tls, error) {
	if _target, ok := (*args)["target"]; ok {
		target := _target.(string)
		m := reTarget.FindStringSubmatch(target)
		if len(m) == 0 {
			return nil, nil, errors.New("target must be provided in the form of: tcp://target:port, udp://target:port, or target:port (defaults to tcp)")
		}

		proto := "tcp"

		var port int64 = 443
		if len(m[2]) != 0 {
			rawPort, err := strconv.ParseUint(m[2][1:], 10, 64)
			if err != nil {
				return nil, nil, errors.New("failed to parse port: " + m[2])
			}
			port = int64(rawPort)
		}

		socket, err := s.Runtime.CreateResource("socket",
			"protocol", proto,
			"port", port,
			"address", m[1],
		)
		if err != nil {
			return nil, nil, err
		}

		(*args)["socket"] = socket
		delete(*args, "target")
	}

	return args, nil, nil
}

func (s *lumiTls) id() (string, error) {
	socket, err := s.Socket()
	if err != nil {
		return "", err
	}

	return "tls+" + socket.LumiResource().Id, nil
}

func (s *lumiTls) GetParams(socket Socket) (map[string]interface{}, error) {
	host, err := socket.Address()
	if err != nil {
		return nil, err
	}

	port, err := socket.Port()
	if err != nil {
		return nil, err
	}

	proto, err := socket.Protocol()
	if err != nil {
		return nil, err
	}

	tester := tlsshake.New(proto, host, int(port))
	if err := tester.Test(); err != nil {
		return nil, err
	}

	res := map[string]interface{}{}
	findings := tester.Findings

	lists := map[string][]string{
		"errors": findings.Errors,
	}
	for field, data := range lists {
		v := make([]interface{}, len(data))
		for i := range data {
			v[i] = data[i]
		}
		res[field] = v
	}

	maps := map[string]map[string]bool{
		"versions": findings.Versions,
		"ciphers":  findings.Ciphers,
	}
	for field, data := range maps {
		v := make(map[string]interface{}, len(data))
		for k, vv := range data {
			v[k] = vv
		}
		res[field] = v
	}

	certs := []interface{}{}
	for i := range findings.Certificates {
		cert := findings.Certificates[i]

		raw, err := s.Runtime.CreateResource("certificate",
			"pem", "",
			// NOTE: if we do not set the hash here, it will generate the cache content before we can store it
			// we are using the hashs for the id, therefore it is required during creation
			"fingerprints", certFingerprints(cert),
		)
		if err != nil {
			return nil, err
		}

		// store parsed object with resource
		lumiCert := raw.(Certificate)
		lumiCert.LumiResource().Cache.Store("_cert", &lumi.CacheEntry{Data: cert})
		certs = append(certs, lumiCert)
	}
	res["certificates"] = certs

	return res, nil
}

func (s *lumiTls) GetVersions(params map[string]interface{}) ([]interface{}, error) {
	raw, ok := params["versions"]
	if !ok {
		return []interface{}{}, nil
	}

	data := raw.(map[string]interface{})
	res := []interface{}{}
	for k, v := range data {
		if v.(bool) {
			res = append(res, k)
		}
	}

	return res, nil
}

func (s *lumiTls) GetCiphers(params map[string]interface{}) ([]interface{}, error) {
	raw, ok := params["ciphers"]
	if !ok {
		return []interface{}{}, nil
	}

	data := raw.(map[string]interface{})
	res := []interface{}{}
	for k, v := range data {
		if v.(bool) {
			res = append(res, k)
		}
	}

	return res, nil
}

func (s *lumiTls) GetCertificates(params map[string]interface{}) ([]interface{}, error) {
	raw, ok := params["certificates"]
	if !ok {
		return []interface{}{}, nil
	}

	return raw.([]interface{}), nil
}