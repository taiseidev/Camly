import 'package:camly_di/camly_di.dart';
import 'package:camly_domain/camly_domain.dart';
import 'package:camly_system/camly_system.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../../constants.dart';

final class AuthRepository extends IAuthRepository {
  AuthRepository(this._ref);

  final Ref _ref;

  @override
  Future<Result<void>> signUp({
    required String email,
    required String password,
  }) async {
    final baseUrl = _ref.read(baseUrlProvider);

    try {
      await ApiClient(baseUrl).post(
        Endpoint.signUp,
        data: {
          Field.email: email,
          Field.password: password,
        },
      );

      return const Success(null);
    } on ApiException catch (e) {
      return Failure(e);
    } catch (e) {
      return Failure(ApiException(message: 'Unexpected error: $e'));
    }
  }
}
