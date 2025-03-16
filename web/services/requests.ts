interface ShortenUrlResponse {
  error: boolean;
  message: string;
  data: string;
}

export const shortenUrl = async (url: string): Promise<ShortenUrlResponse> => {
  const apiUrl = process.env.NEXT_PUBLIC_API_URL;

  const response = await fetch(`${apiUrl}/create`, {
    method: "POST",
    body: JSON.stringify({ link: url }),
    headers: {
      "Content-Type": "application/json",
    },
  });

  const data: ShortenUrlResponse = await response.json();
  return data;
};
