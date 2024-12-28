{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = with pkgs; [
    go
  ];

  name = "jlpt-notify";

  shellHook = ''
    exec zsh;
  '';
}
