import 'package:camly_di/src/providers/auth_repository_provider.dart';
import 'package:camly_domain/camly_domain.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

part 'sign_up_usecase_provider.g.dart';

@Riverpod(keepAlive: true)
SignUpUseCase signUpUseCase(Ref ref) {
  final authRepository = ref.read(authRepositoryProvider);
  return SignUpUseCase(authRepository);
}
