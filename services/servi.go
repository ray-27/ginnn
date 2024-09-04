package services

func FetchData(key string) (string, error) {
    result, err := client.Get(ctx, key).Result()
    if err != nil {
        return "", err
    }
    return result, nil
}