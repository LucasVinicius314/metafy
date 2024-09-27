class MainEvent {
  const MainEvent();
}

class ConnectMainEvent extends MainEvent {
  const ConnectMainEvent({
    required this.token,
  });

  final String token;
}
