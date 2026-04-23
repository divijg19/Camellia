import 'package:flutter/material.dart';
import 'package:shared_core/core.dart';

void main() {
  runApp(const CamelliaApp());
}

class CamelliaApp extends StatefulWidget {
  const CamelliaApp({super.key});

  @override
  State<CamelliaApp> createState() => _CamelliaAppState();
}

class _CamelliaAppState extends State<CamelliaApp> {
  final _api = BackendApi(baseUrl: 'http://10.0.2.2:3000');
  late Future<List<ClassItem>> _classesFuture;

  @override
  void initState() {
    super.initState();
    _classesFuture = _api.fetchClasses();
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(title: const Text('Camellia Flutter Shell')),
        body: FutureBuilder<List<ClassItem>>(
          future: _classesFuture,
          builder: (context, snapshot) {
            if (snapshot.connectionState == ConnectionState.waiting) {
              return const Center(child: CircularProgressIndicator());
            }

            if (snapshot.hasError) {
              return Center(child: Text('Error: ${snapshot.error}'));
            }

            final classes = snapshot.data ?? const <ClassItem>[];
            if (classes.isEmpty) {
              return const Center(child: Text('No classes available'));
            }

            return ListView.builder(
              itemCount: classes.length,
              itemBuilder: (context, index) {
                final item = classes[index];
                return ListTile(
                  title: Text(item.title),
                  subtitle: Text('${item.instructor} • ${item.location}'),
                  trailing: Text('\$${item.price.toStringAsFixed(2)}'),
                );
              },
            );
          },
        ),
      ),
    );
  }
}
