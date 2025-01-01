import 'package:camly_data/camly_data.dart';
import 'package:camly_domain/camly_domain.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

import 'use_mock_provider.dart';

part 'auth_repository_provider.g.dart';

@Riverpod(keepAlive: true)
IAuthRepository authRepository(Ref ref) {
  final useMock = ref.read(useMockProvider);
  return switch (useMock) {
    true => MockAuthRepository(),
    false => AuthRepository(ref),
  };
}
