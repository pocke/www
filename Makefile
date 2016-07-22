release:
	gox --arch 'amd64 386' --os 'windows linux darwin' --output "dist/{{.Dir}}_{{.OS}}_{{.Arch}}/{{.Dir}}"
	zip      pkg/www_windows_386.zip     dist/www_windows_386/www.exe   -j
	zip      pkg/www_windows_amd64.zip   dist/www_windows_amd64/www.exe -j
	tar zcvf pkg/www_linux_386.tar.gz    -C dist/www_linux_386/    www
	tar zcvf pkg/www_linux_amd64.tar.gz  -C dist/www_linux_amd64/  www
	tar zcvf pkg/www_darwin_386.tar.gz   -C dist/www_darwin_386/   www
	tar zcvf pkg/www_darwin_amd64.tar.gz -C dist/www_darwin_amd64/ www

clean:
	rm -rf dist/
	rm -f pkg/*.tar.gz pkg/*.zip
