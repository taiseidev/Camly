import 'package:flutter_secure_storage/flutter_secure_storage.dart';

class SecureStorageSingleton {
  static final SecureStorageSingleton _instance =
      SecureStorageSingleton._internal();

  final FlutterSecureStorage _secureStorage;

  SecureStorageSingleton._internal()
      : _secureStorage = const FlutterSecureStorage();

  factory SecureStorageSingleton() {
    return _instance;
  }

  FlutterSecureStorage get instance => _secureStorage;
}
