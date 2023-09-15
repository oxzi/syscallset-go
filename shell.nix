# SPDX-FileCopyrightText: Alvar Penning
#
# SPDX-License-Identifier: BSD-3-Clause

{ config ? {}, overlays ? [], pkgs ? import <nixpkgs> { } }:

pkgs.mkShell {
  buildInputs = with pkgs; [
    audit
    cowsay
    go_1_17
    golangci-lint
    reuse
    shellcheck
  ];
}
