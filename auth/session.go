func (s *SessionStore) Logout(ctx context.Context, refreshToken string) error {
    // Invalidate session in DB
    if err := s.db.Invalidate(ctx, refreshToken); err != nil {
        return err
    }
    // Invalidate in cache immediately (was up to 60s delay)
    return s.cache.Delete(ctx, "session:"+hashToken(refreshToken))
}
