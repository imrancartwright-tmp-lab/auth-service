func (v *Validator) validateAudience Claims {
    if v.StrictAudience {
        // Only allow tokens intended for this specific service
        for _, aud := range claims.Audience {
            if aud == v.ServiceAudience {
                return nil
            }
        }
        return ErrAudienceMismatch
    }
    return nil
}
