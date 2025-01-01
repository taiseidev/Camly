import 'package:camly_domain/camly_domain.dart';

final class SignUpUseCase {
  final IAuthRepository authRepository;

  SignUpUseCase(this.authRepository);

  Future<void> execute({
    required String email,
    required String password,
  }) async {
    if (email.isEmpty || password.isEmpty) {
      throw ArgumentError('Email and password cannot be empty.');
    }

    await authRepository.signUp(
      email: email,
      password: password,
    );
  }
}
