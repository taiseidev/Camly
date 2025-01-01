// ignore_for_file: do_not_use_environment

import 'package:camly_di/camly_di.dart';
import 'package:camly_i18n/gen/strings.g.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'app.dart';

void main() {
  WidgetsFlutterBinding.ensureInitialized();

  // デバイスのロケールで初期化
  LocaleSettings.useDeviceLocale();

  final environment = switch (const String.fromEnvironment('FLAVOR')) {
    'dev' => Environment.dev,
    'stg' => Environment.stg,
    'prod' => Environment.prod,
    _ => throw UnimplementedError(),
  };

  const useMock = bool.fromEnvironment('MOCK');

  runApp(
    ProviderScope(
      overrides: [
        environmentProvider.overrideWith((_) => environment),
        useMockProvider.overrideWith((_) => useMock),
      ],
      child: TranslationProvider(
        child: const App(),
      ),
    ),
  );
}
