import 'package:camly_system/camly_system.dart';
import 'package:dio/dio.dart';

final class UpdateTokensInterceptor extends Interceptor {
  const UpdateTokensInterceptor();

  @override
  Future<void> onResponse(
    Response response,
    ResponseInterceptorHandler handler,
  ) async {
    if (response.data is Map<String, dynamic>) {
      final data = response.data as Map<String, dynamic>;

      if (data.containsKey('accessToken')) {
        await SecureStorageKey.accessToken.save(data['accessToken'] as String);
      }

      if (data.containsKey('refreshToken')) {
        await SecureStorageKey.refreshToken
            .save(data['refreshToken'] as String);
      }

      if (data.containsKey('accessTokenExpiration')) {
        await SecureStorageKey.accessTokenExpiration
            .save(data['accessTokenExpiration'] as String);
      }

      if (data.containsKey('refreshTokenExpiration')) {
        await SecureStorageKey.refreshTokenExpiration
            .save(data['refreshTokenExpiration'] as String);
      }
    }

    return handler.next(response);
  }
}
