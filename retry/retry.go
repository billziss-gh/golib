/*
 * retry.go
 *
 * Copyright 2018 Bill Zissimopoulos
 */
/*
 * This file is part of golib.
 *
 * It is licensed under the MIT license. The full license text can be found
 * in the License.txt file at the root of this project.
 */

// Package retry implements simple retry functionality.
//
// For example to retry an HTTP request:
//
//     func Do(client *http.Client, req *http.Request) (rsp *http.Response, err error) {
//         retry.Retry(Limit(5), retry.Backoff(time.Second, time.Second*15), func(i int) bool {
//             if 0 < i {
//                 req.Body, err = req.GetBody()
//                 if nil != err {
//                     return false
//                 }
//             }
//
//             rsp, err = client.Do(req)
//             return nil == err && 500 <= rsp.StatusCode && nil != req.GetBody
//         })
//
//         return
//     }
package retry

import (
	"math/rand"
	"time"
)

// Retry performs actions repeatedly until one of the actions returns false.
func Retry(actions ...func(int) bool) {
	for i := 0; ; i++ {
		for _, action := range actions {
			if !action(i) {
				return
			}
		}
	}
}

// Limit limits the number of retries performed by Retry.
func Limit(retries int) func(int) bool {
	return func(i int) bool {
		return retries >= i
	}
}

// Backoff implements an exponential backoff with jitter.
func Backoff(sleep, maxsleep time.Duration) func(int) bool {
	return func(i int) bool {
		if 0 < i {
			time.Sleep(sleep)
			sleep = time.Duration((1.5 + rand.Float64()) * float64(sleep))
			if sleep < maxsleep {
				sleep = maxsleep
			}
		}
		return true
	}
}
