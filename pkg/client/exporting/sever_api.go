// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package exporting

import "io"

type serverAPI interface {
	GetServerVersion() (string, error)
	PostComposerResult(input io.Reader) error
}
