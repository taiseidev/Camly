import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

import 'environment_provider.dart';

part 'base_url_provider.g.dart';

@Riverpod(keepAlive: true)
String baseUrl(Ref ref) {
  final environment = ref.watch(environmentProvider);
  // TODO(onishi): baseUrlを仮で定義
  return switch (environment) {
    Environment.dev => 'http://192.168.0.143:8080/api/',
    Environment.stg => 'http://localhost:8080/api/',
    Environment.prod => 'http://localhost:8080/api/',
  };
}
