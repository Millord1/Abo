module default {
    type Subscription {
        required property name -> str { constraint exclusive };
        required property blockchain -> str;
        required property lastMint -> str;
        required multi link types -> Type
    }

    type Type {
        required property name -> str { constraint exclusive };
        required link subscription -> Subscription;
        multi link assets -> Asset {
            property quantity -> int32
        }
    }

    type Contract {
        required property address -> str { constraint exclusive };
        required property standard -> str;
        required property blockchain -> str;
    }

    type Asset {
        required property assetId -> str { constraint exclusive };
        required property name -> str;
        required property imgUrl -> str;
        required property metadata -> str;
        required property description -> str;
        multi link contracts -> Contract;
        link collection -> Collection;
    }

    type Collection {
        required property collectionId -> str { constraint exclusive };
        required property name -> str;
        required property imgUrl -> str;
        required property description -> str;
    }

    type User {
        required property email -> str { constraint exclusive };
        required property ethAddress -> str;
        required property userName -> str;
        property foAddress -> str;
        property xcpAddress -> str;
        link hasType -> Type;
        multi link mintLogs -> MintLog
    }

    type MintLog {
        required property txHash -> str { constraint exclusive };
        required property date -> str;
        required link contract -> Contract;
        required link asset -> Asset
    }
}
