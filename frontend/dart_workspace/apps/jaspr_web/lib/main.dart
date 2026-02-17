import 'package:jaspr/jaspr.dart';
import 'package:jaspr/dom.dart';
import 'package:shared_core/shared_core.dart';

void main() {
  runApp(const App());
}

class App extends StatefulComponent {
  const App({super.key});

  @override
  State<App> createState() => _AppState();
}

class _AppState extends State<App> {
  final _api = BackendApi(baseUrl: 'http://localhost:3000');
  List<ClassItem> _classes = const [];
  String? _error;

  @override
  void initState() {
    super.initState();
    loadClasses();
  }

  Future<void> loadClasses() async {
    try {
      final classes = await _api.fetchClasses();
      setState(() {
        _classes = classes;
        _error = null;
      });
    } catch (err) {
      setState(() {
        _error = err.toString();
      });
    }
  }

  @override
  Component build(BuildContext context) {
    final classList = <Component>[
      for (final classItem in _classes)
        li([
          Component.text(
              '${classItem.title} · ${classItem.instructor} · \$${classItem.price}')
        ]),
    ];

    return div(classes: 'p-6', [
      h1([Component.text('Camellia (Jaspr Web)')]),
      p([
        Component.text(
            'Data is loaded from Go backend /api/classes via shared_core.')
      ]),
      button(
        classes: 'mt-2 rounded bg-slate-900 px-3 py-2 text-white',
        onClick: () => loadClasses(),
        [Component.text('Reload classes')],
      ),
      if (_error != null)
        p(classes: 'mt-3 text-red-600', [Component.text(_error!)]),
      ul(classList),
    ]);
  }
}
