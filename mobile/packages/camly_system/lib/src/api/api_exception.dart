import 'package:dio/dio.dart';

/// カスタム例外クラス
final class ApiException implements Exception {
  final String message;
  final int? statusCode;

  ApiException({
    required this.message,
    this.statusCode,
  });

  factory ApiException.fromDioError(DioException error) {
    return ApiException(
      // ignore: avoid_dynamic_calls
      message: error.response?.data['message'] ?? error.message,
      statusCode: error.response?.statusCode,
    );
  }

  @override
  String toString() {
    return 'ApiException: $message (Status code: $statusCode)';
  }
}
