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

studio:
	maestro studio

maestro_ios_test:
	cd apps/camly_app && \
	flutter build ios --debug --simulator --dart-define=FLAVOR=dev -t lib/main.dart && \
	cd ../../
	maestro test .maestro/camly_app/test_ios.yml

maestro_android_test:
	cd apps/camly_app && \
	flutter build apk --debug --dart-define=FLAVOR=dev -t lib/main.dart && \
	cd ../../
	maestro test .maestro/camly_app/test_android.yml
