import 'package:camly_domain/src/result.dart';

abstract class IAuthRepository {
  Future<Result<void>> signUp({
    required String email,
    required String password,
  });
}
