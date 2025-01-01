import 'secure_storage_singleton.dart';

enum SecureStorageKey {
  // トークン
  accessToken,
  refreshToken,

  // トークンの有効期限
  accessTokenExpiration,
  refreshTokenExpiration;

  Future<void> save(String value) async {
    await SecureStorageSingleton().instance.write(key: name, value: value);
  }

  Future<String?> read() async {
    return SecureStorageSingleton().instance.read(key: name);
  }

  Future<void> delete() async {
    await SecureStorageSingleton().instance.delete(key: name);
  }
}
