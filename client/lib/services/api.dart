import 'dart:convert';

import 'package:http/http.dart' as http;

class Api {
  Future<dynamic> post({
    required String path,
    required Map<String, dynamic> body,
  }) async {
    return await http.post(Uri(), body: jsonEncode(body));
  }
}
