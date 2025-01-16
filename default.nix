with import <nixpkgs> {};

stdenv.mkDerivation {
  name = "acre-webhook";
  buildInputs = with pkgs; [
    go
    gnumake
  ];
  hardeningDisable = [ "fortify" ];
  shellHook = ''
    export GOPRIVATE=github.com/activatedio/*
  '';
}
