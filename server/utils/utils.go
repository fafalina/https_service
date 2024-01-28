package utils

import (
    "crypto/x509"
    "io/ioutil"
    "log"
)

/*
 * @brief Load CA certificate and put it into new pool
 */
func LoadCA(caFile string) *x509.CertPool {
    pool := x509.NewCertPool()

    if ca, e := ioutil.ReadFile(caFile); e != nil {
        log.Fatal("ReadFile: ", e)
    } else {
        pool.AppendCertsFromPEM(ca)
    }
    return pool
}
