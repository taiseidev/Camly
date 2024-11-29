import 'shared_preference_singleton.dart';

enum SharedPreferenceKey {
  isFirst;

  final prefs = SharedPreferencesSingleton();

  Future<void> save(Object value) async {
    if (value is String) {
      await prefs.instance.setString(name, value);
    }

    if (value is int) {
      await prefs.instance.setInt(name, value);
    }

    if (value is bool) {
      await prefs.instance.setBool(name, value);
    }

    if (value is double) {
      await prefs.instance.setDouble(name, value);
    }

    if (value is List<String>) {
      await prefs.instance.setStringList(name, value);
    }
  }

  Future<String?> getString() async => prefs.instance.getString(name);

  Future<int?> getInt() async => prefs.instance.getInt(name);

  Future<bool?> getBool() async => prefs.instance.getBool(name);

  Future<double?> getDouble() async => prefs.instance.getDouble(name);

  Future<List<String>?> getStringList() async =>
      prefs.instance.getStringList(name);

  Future<void> delete() async => prefs.instance.remove(name);

  Future<void> deleteAll() async => prefs.instance.clear();
}
