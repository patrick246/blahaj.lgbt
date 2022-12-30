import {useEffect, useState} from "react";

export interface BlahajAvailability {
    store_id: string;
    store_name: string;
    store_country: string;
    location: GeoCoordinates;
    number: number;
}

export interface GeoCoordinates {
    latitude: string;
    longitude: string;
}

export interface CountryAvailability {
    country: string;
    number: number;
}

export function useCountryAvailability(): [CountryAvailability[] | undefined, boolean] {
    const [availability, setAvailability] = useState<CountryAvailability[]>();
    const [isLoading, setIsLoading] = useState<boolean>(false);

    useEffect(() => {
        setIsLoading(true);
        fetch('/api/availability/countries')
            .then(res => res.json())
            .then((a: CountryAvailability[]) => {
                a.sort((a1, a2) => {
                    return a1.country.localeCompare(a2.country)
                });

                setAvailability(a);
                setIsLoading(false);
            });
    }, []);

    return [availability, isLoading];
}

export function useStoreAvailability(country: string = ""): [BlahajAvailability[] | undefined, boolean] {
    const [availability, setAvailability] = useState<BlahajAvailability[]>();
    const [isLoading, setIsLoading] = useState<boolean>(false);

    useEffect(() => {
        setIsLoading(true);
        fetch('/api/availability/stores')
            .then(res => res.json())
            .then((a: BlahajAvailability[]) => {
                a.sort((a1, a2) => {
                    return a1.store_name.localeCompare(a2.store_name)
                });

                if (country !== "") {
                    a = a.filter(availabilityEntry => availabilityEntry.store_country == country)
                }

                setAvailability(a);
                setIsLoading(false);
            });
    }, [country]);

    return [availability, isLoading];

}