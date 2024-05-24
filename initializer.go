/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package reconciler

import (
	"context"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type initializer[T client.Object] struct{ Funcs[T] }

func (t *initializer[T]) Reconcile(ctx context.Context, resource T) (ctrl.Result, error) {
	log := logr.FromContextOrDiscard(ctx).WithCallDepth(resultCallDepth)
	log.Info("Reconciler has been triggered")

	return t.Next(ctx, resource)
}
