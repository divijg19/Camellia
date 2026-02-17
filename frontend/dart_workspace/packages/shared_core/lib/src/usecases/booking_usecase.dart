import '../models/class_item.dart';
import '../services/backend_api.dart';

class BookingUseCase {
  final BackendApi backendApi;

  const BookingUseCase(this.backendApi);

  Future<String> bookClass({
    required String userId,
    required ClassItem classItem,
  }) async {
    final message = await backendApi.createBooking(
      userId: userId,
      classId: classItem.id,
    );
    return '$message for ${classItem.title}';
  }
}
