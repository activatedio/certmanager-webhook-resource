with import <nixpkgs> {};

stdenv.mkDerivation {
  name = "acre-bluebird-operator";
  buildInputs = with pkgs; [
    go
    gnumake
  ];
  hardeningDisable = [ "fortify" ];
}
