package cryptox

import (
	"testing"
)

// 私钥生成
// openssl genrsa -out rsa_private_key.pem 1024
var privateKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCeHtkgsdNNmzHA5memsqcDcx+8/Bfx3pmG5CfIEATJzlERLoebadyg1/3pG6M8N5W3ngnoMd7Q8nK4f2goKmMpq06eytkUfmd31EB8o2a65Jel8Ask1bmx+c3tUKPGk0y8PT2jeEoYhoZx12SEqU2KR0yu8oT7zS9xkymIJ3KBd19lNgrw2iM3wDf8Kdef0bfuFS3SqULUtZpMbLbkVZpA7CKviXO+/qC2iNohjaAK3UX7R3ivmxBmSEIMevK+9Axrg28+1TkRI4W0wK8JR+iQNHBcMlOAY/fGjDfIdYq0VI/CKTjUtej2+QaF0dxABrmFVt5YmPuXx3B/oZYySqpLAgMBAAECggEBAJ0ZlS0wwPqISWSZBz7OnbWFunDwH3JGck9dfFD/6WIARPdHPaCiC30bXMEwnnyD1ZL7yGgUtIbxfB3nKDtP8fpCo/iNHiWIeZDLJ88uIjFqUWmjF89fMgKsaXdx+Wagj0svpYA3UZcQgZ+2G77a35mVwZDIkGwFry+uwULvvor10tS9CsCiEM/BP/dF6XJkxXZwGtqllxHrSvqZjFW4eEiupIduUvYqhiSQS0XQQnxtAKOjq+jcDUiXDKI2pe1STtHgs4Xk9duf99uOBaoJsiZ5IRbnJq1yPNO5c5bgjM0Oi6h9L7cH0gBuNcNgeRMNHEj8e9c7j9FRu8ViVvABbbECgYEA8qA9Rsc/m0leRoQuZmMD9J8yrVYQLVwb2ZRF6msgg4GVyRLIa/AxJ4JTb3hA00DkG9hVUFb0wf24l+pdJSN0pJ4FcFXHLFzliXjTi/0+/nKeLE1ASQ+yD4630O7zf6VmZCSaZw759us0JPNIi8VjJG3O1uFT6s/Oc8d8rvqVaI8CgYEAptYhLUk308eX8MJKDhsV9ngCulEtD6+N+mlcY+GGPoXAf2DvkiJVr26DFt0Ole7ce8rouT76BkUAt0SBNVSCJi8jcDtDFF5KeVahRsbfi7vV2U7xS6GPi0yyhbADgMlkMXQltk49bnacy0w7LzxQqRBbbyBuhA8et9N10kjbKIUCgYBQX4CtB6gsett+JJ9yT26qBGq3PQFS5WZcZ1/Zze8RYUYLO+ukxgaifevBy2Y+FQr7pmhsmsGt2azitkJw4RPszGQk2GWIbqWjKqF5iETr6XA4YMbg7RAhUzfVQLLtIA+RMQbtCOsJPtUhQG5RzuuOz8lrKW2xJXHQrM7C8EO6RQKBgGYPwKQPkpOAGrcyZFNUxo7Mw/5vAn4CC30RT2/a22EMOcKy7lwbnQeBN7iY6v9V7JPegZgxqoqDU/jJc7HITQZG+AsEdfjT5gNst6SwrLiy5BPYc0ytPT8eYGfgUYRLr3uBbkCkhHg7H5hOvjYbU8zZMNY+pN0s0sHqHr3RzCBxAoGAcpABj4uRs8iBPjYVxMu9Q6L+SIrHIgBLdYfMQLiyO7V2FigMfHZiBp1bNTox1IjxasjL8ALcbYmoa67nYOx5VQTgsgP9pj6jo8dIqEZogY0qFH+qvMbjGL3U98Ic4pbixDmoGYopF8im10eIrFJ8lxGe7hGMubQa1YH9k5vZBNk=
-----END RSA PRIVATE KEY-----
`

// 公钥: 根据私钥生成
// openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var publicKey = `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDcGsUIIAINHfRTdMmgGwLrjzfM
NSrtgIf4EGsNaYwmC1GjF/bMh0Mcm10oLhNrKNYCTTQVGGIxuc5heKd1gOzb7bdT
nCDPPZ7oV7p1B9Pud+6zPacoqDz2M24vHFWYY2FbIIJh8fHhKcfXNXOLovdVBE7Z
y682X1+R1lRK8D+vmQIDAQAB
-----END PUBLIC KEY-----
`

func TestRsa(t *testing.T) {
	data := "hello world"
	ed, err := EncryptRSA(publicKey, data)
	t.Logf("data: %v, 加密后: %v, err: %v", string(data), ed, err)

	dd, err := DecryptRSA(privateKey, "eVbiTRetzkOr/IQycNmIqJcTrtmEpg2gMFUjHfsCT9f0uRWuRV6afdUK6Kpc4rApi9wpf9UfFpqNmlRO6sc3KQPRlC9WwC398kUex4NydVlsYvf8Dn/ZbHtMsFVl1Js4c/aw2d17coJ8sn+jYlphSz79Ef7811pP2iVT/7ldmGU/qe/BZC/U1YuzhdtPh6L9xe8PGtLJF1o2GQMeFnNkLa0fY8KC1osGQJF7yYpt6AJhN8CvHSEiu3KRFP1NKBbhRgwVRRflF1moBHUf5w9YO+CcuCICQuFYPKk/PYMeWvGgDMjHOOOCgoXRuLrUTRa4mXhf9vj8GqjwEfnL/4V8tg==")

	t.Logf("解密: %v, err: %v", dd, err)
}
