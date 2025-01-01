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

FLUTTER_BUILD_COMMON_ARGS := --debug --dart-define=FLAVOR=dev -t lib/main.dart

maestro_ios_test:
	cd apps/camly_app && \
	flutter build ios \
		$(FLUTTER_BUILD_COMMON_ARGS) \
		--simulator && \
	cd ../../
	maestro test .maestro/camly_app/test_ios.yml

maestro_android_test:
	cd apps/camly_app && \
	flutter build apk \
		$(FLUTTER_BUILD_COMMON_ARGS) && \
	cd ../../
	maestro test .maestro/camly_app/test_android.yml
