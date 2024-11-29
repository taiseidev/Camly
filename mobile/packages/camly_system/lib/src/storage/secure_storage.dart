import 'secure_storage_singleton.dart';

enum SecureStorageKey {
  accessToken;

  final storage = SecureStorageSingleton();

  Future<void> save(String value) async {
    await storage.instance.write(key: name, value: value);
  }

  Future<String?> read() async {
    return storage.instance.read(key: name);
  }

  Future<void> delete() async {
    await storage.instance.delete(key: name);
  }
}
