class ClassItem {
  final String id;
  final String title;
  final String instructor;
  final String location;
  final double price;

  const ClassItem({
    required this.id,
    required this.title,
    required this.instructor,
    required this.location,
    required this.price,
  });

  factory ClassItem.fromJson(Map<String, dynamic> json) {
    return ClassItem(
      id: (json['id'] ?? '').toString(),
      title: (json['title'] ?? json['name'] ?? '').toString(),
      instructor: (json['instructor'] ?? '').toString(),
      location: (json['location'] ?? 'Online').toString(),
      price: (json['price'] as num?)?.toDouble() ?? 0,
    );
  }

  String get label => '$title by $instructor ($location)';
}
