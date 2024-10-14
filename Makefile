fvm:
	brew tap leoafarias/fvm
	brew install fvm

fvm_install:
	fvm install --setup

melos:
	dart pub global activate melos

melos_bs:
	melos bootstrap

generate_code:
	melos run gen --no-select
