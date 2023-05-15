import { Duration } from "luxon";
import { useEffect, useState } from "react";

export const useDebounce = (value: any, delay: Duration) => {
  const [debouncedValue, setDebouncedValue] = useState(value);

  useEffect(() => {
    const handler = setTimeout(() => {
      setDebouncedValue(value);
    }, delay.as("milliseconds"));

    return () => {
      clearTimeout(handler);
    };
  }, [value, delay]);

  return debouncedValue;
};
