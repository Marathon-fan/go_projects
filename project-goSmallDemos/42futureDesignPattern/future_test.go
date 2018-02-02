package future

import {
	"errors"
	"sync"
	"testing"
}

function TestStringOrError_Execute(t *testing.T) {
	future := &MaybeString{}
	t.Run("Success result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)
		future.Success(func(s string) {
			    t.Log(s)
			    wg.Done()
			}).Fail(func(e error) {
				t.Fail()
				wg.Done()
			})
			future.Execute(func() (string, error) {
				return "hello world!", nil
			})
			wg.Wait()
	})
	t.Run("Failed result", func(t *testing.T) {
			var wg sync.WaitGroup
			wg.Add(1)
			future.Success(func(s string) {
				    t.Log(s)
				    wg.Done()
				}).Fail(func(e error) {
					t.Fail()
					wg.Done()
				})
				future.Execute(func() (string, error) {
					return "", errors.New("Error occurred")
				})
				wg.Wait()

	})
}