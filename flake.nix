{
  outputs = { self, nixpkgs }: {

    overlay = final: prev: {
      py-go-embed = final.buildGoModule {
        name = "py-go-embed";
        src = self;
        vendorSha256 = "sha256-pQpattmS9VmO3ZIQUFn66az8GSmB4IvYhTTCFn6SUmo=";
        nativeBuildInputs = with final; [ pkg-config ];
        buildInputs = with final; [ python3 ncurses ];
      };
    };

    packages.x86_64-linux.py-go-embed = (import nixpkgs {
      system = "x86_64-linux";
      overlays = [ self.overlay ];
    }).py-go-embed;

    defaultPackage.x86_64-linux = self.packages.x86_64-linux.py-go-embed;
  };
}
