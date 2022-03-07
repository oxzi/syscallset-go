# SPDX-FileCopyrightText: 2022 Alvar Penning
#
# SPDX-License-Identifier: BSD-3-Clause

{ config ? {}, overlays ? [], pkgs ? import <nixpkgs> { } }:

pkgs.mkShell {
  buildInputs = with pkgs; [
    go_1_17
    golangci-lint
    reuse
  ];
}
