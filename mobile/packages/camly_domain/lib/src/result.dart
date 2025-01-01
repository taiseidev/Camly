sealed class Result<S> {
  const Result();
}

final class Success<S> extends Result<S> {
  const Success(this.value);
  final S value;
}

final class Failure<E> extends Result<Never> {
  const Failure(this.exception);
  final E exception;
}
