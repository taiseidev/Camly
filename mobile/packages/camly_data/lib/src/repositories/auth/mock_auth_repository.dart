import 'package:camly_domain/camly_domain.dart';

final class MockAuthRepository extends IAuthRepository {
  @override
  Future<Result<void>> signUp({
    required String email,
    required String password,
  }) async {
    return const Success(null);
  }
}
