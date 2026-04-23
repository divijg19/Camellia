import 'dart:convert';

import 'package:http/http.dart' as http;

import '../models/class_item.dart';

class BackendApi {
  final String baseUrl;
  final http.Client _client;

  BackendApi({required this.baseUrl, http.Client? client})
      : _client = client ?? http.Client();

  Future<List<ClassItem>> fetchClasses() async {
    final uri = Uri.parse('$baseUrl/api/classes');
    final response = await _client.get(uri);

    if (response.statusCode < 200 || response.statusCode >= 300) {
      throw BackendApiException(
        'Failed to fetch classes (${response.statusCode})',
      );
    }

    final decoded = jsonDecode(response.body);
    if (decoded is! List) {
      throw BackendApiException('Invalid classes response format');
    }

    return decoded
        .whereType<Map>()
        .map((item) => ClassItem.fromJson(Map<String, dynamic>.from(item)))
        .toList(growable: false);
  }

  Future<String> createBooking({
    required String userId,
    required String classId,
  }) async {
    final uri = Uri.parse('$baseUrl/api/bookings');
    final response = await _client.post(
      uri,
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({'userId': userId, 'classId': classId}),
    );

    if (response.statusCode < 200 || response.statusCode >= 300) {
      throw BackendApiException(
        'Failed to create booking (${response.statusCode})',
      );
    }

    final decoded = jsonDecode(response.body);
    if (decoded is! Map<String, dynamic>) {
      throw BackendApiException('Invalid booking response format');
    }

    return (decoded['message'] as String?) ?? 'Booking successful';
  }
}

class BackendApiException implements Exception {
  final String message;

  BackendApiException(this.message);

  @override
  String toString() => 'BackendApiException: $message';
}
