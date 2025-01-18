with import <nixpkgs> {};

stdenv.mkDerivation {
  name = "certmanager-webhook-resource";
  buildInputs = with pkgs; [
    go
    gnumake
  ];
  hardeningDisable = [ "fortify" ];
  shellHook = ''
    export GOPRIVATE=github.com/activatedio/*
  '';
}
