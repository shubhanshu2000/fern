# This file was auto-generated by Fern from our API Definition.

from seed import SeedApi
from seed import AsyncSeedApi
import typing
from .utilities import validate_response


async def test_get_account(client: SeedApi, async_client: AsyncSeedApi) -> None:
    expected_response: typing.Any = {
        "id": "id",
        "related_resources": [
            {
                "id": "id",
                "related_resources": [
                    {"resource_type": "Account", "name": "name"},
                    {"resource_type": "Account", "name": "name"},
                ],
                "memo": {
                    "description": "description",
                    "account": {"resource_type": "Account", "name": "name"},
                },
                "resource_type": "Account",
                "name": "name",
                "patient": {
                    "id": "id",
                    "related_resources": [],
                    "memo": {"description": "description"},
                    "resource_type": "Patient",
                    "name": "name",
                    "scripts": [
                        {"resource_type": "Script", "name": "name"},
                        {"resource_type": "Script", "name": "name"},
                    ],
                },
                "practitioner": {
                    "id": "id",
                    "related_resources": [],
                    "memo": {"description": "description"},
                    "resource_type": "Practitioner",
                    "name": "name",
                },
            },
            {
                "id": "id",
                "related_resources": [
                    {"resource_type": "Account", "name": "name"},
                    {"resource_type": "Account", "name": "name"},
                ],
                "memo": {
                    "description": "description",
                    "account": {"resource_type": "Account", "name": "name"},
                },
                "resource_type": "Account",
                "name": "name",
                "patient": {
                    "id": "id",
                    "related_resources": [],
                    "memo": {"description": "description"},
                    "resource_type": "Patient",
                    "name": "name",
                    "scripts": [
                        {"resource_type": "Script", "name": "name"},
                        {"resource_type": "Script", "name": "name"},
                    ],
                },
                "practitioner": {
                    "id": "id",
                    "related_resources": [],
                    "memo": {"description": "description"},
                    "resource_type": "Practitioner",
                    "name": "name",
                },
            },
        ],
        "memo": {
            "description": "description",
            "account": {
                "id": "id",
                "related_resources": [
                    {"resource_type": "Account", "name": "name"},
                    {"resource_type": "Account", "name": "name"},
                ],
                "memo": {
                    "description": "description",
                    "account": {"resource_type": "Account", "name": "name"},
                },
                "resource_type": "Account",
                "name": "name",
                "patient": {
                    "id": "id",
                    "related_resources": [],
                    "memo": {"description": "description"},
                    "resource_type": "Patient",
                    "name": "name",
                    "scripts": [
                        {"resource_type": "Script", "name": "name"},
                        {"resource_type": "Script", "name": "name"},
                    ],
                },
                "practitioner": {
                    "id": "id",
                    "related_resources": [],
                    "memo": {"description": "description"},
                    "resource_type": "Practitioner",
                    "name": "name",
                },
            },
        },
        "resource_type": "Account",
        "name": "name",
        "patient": {
            "id": "id",
            "related_resources": [
                {
                    "id": "id",
                    "related_resources": [],
                    "memo": {"description": "description"},
                    "resource_type": "Account",
                    "name": "name",
                    "patient": {
                        "resource_type": "Patient",
                        "name": "name",
                        "scripts": [],
                    },
                    "practitioner": {"resource_type": "Practitioner", "name": "name"},
                },
                {
                    "id": "id",
                    "related_resources": [],
                    "memo": {"description": "description"},
                    "resource_type": "Account",
                    "name": "name",
                    "patient": {
                        "resource_type": "Patient",
                        "name": "name",
                        "scripts": [],
                    },
                    "practitioner": {"resource_type": "Practitioner", "name": "name"},
                },
            ],
            "memo": {
                "description": "description",
                "account": {
                    "id": "id",
                    "related_resources": [],
                    "memo": {"description": "description"},
                    "resource_type": "Account",
                    "name": "name",
                    "patient": {
                        "resource_type": "Patient",
                        "name": "name",
                        "scripts": [],
                    },
                    "practitioner": {"resource_type": "Practitioner", "name": "name"},
                },
            },
            "resource_type": "Patient",
            "name": "name",
            "scripts": [
                {
                    "id": "id",
                    "related_resources": [
                        {"resource_type": "Account", "name": "name"},
                        {"resource_type": "Account", "name": "name"},
                    ],
                    "memo": {
                        "description": "description",
                        "account": {"resource_type": "Account", "name": "name"},
                    },
                    "resource_type": "Script",
                    "name": "name",
                },
                {
                    "id": "id",
                    "related_resources": [
                        {"resource_type": "Account", "name": "name"},
                        {"resource_type": "Account", "name": "name"},
                    ],
                    "memo": {
                        "description": "description",
                        "account": {"resource_type": "Account", "name": "name"},
                    },
                    "resource_type": "Script",
                    "name": "name",
                },
            ],
        },
        "practitioner": {
            "id": "id",
            "related_resources": [
                {
                    "id": "id",
                    "related_resources": [],
                    "memo": {"description": "description"},
                    "resource_type": "Account",
                    "name": "name",
                    "patient": {
                        "resource_type": "Patient",
                        "name": "name",
                        "scripts": [],
                    },
                    "practitioner": {"resource_type": "Practitioner", "name": "name"},
                },
                {
                    "id": "id",
                    "related_resources": [],
                    "memo": {"description": "description"},
                    "resource_type": "Account",
                    "name": "name",
                    "patient": {
                        "resource_type": "Patient",
                        "name": "name",
                        "scripts": [],
                    },
                    "practitioner": {"resource_type": "Practitioner", "name": "name"},
                },
            ],
            "memo": {
                "description": "description",
                "account": {
                    "id": "id",
                    "related_resources": [],
                    "memo": {"description": "description"},
                    "resource_type": "Account",
                    "name": "name",
                    "patient": {
                        "resource_type": "Patient",
                        "name": "name",
                        "scripts": [],
                    },
                    "practitioner": {"resource_type": "Practitioner", "name": "name"},
                },
            },
            "resource_type": "Practitioner",
            "name": "name",
        },
    }
    expected_types: typing.Any = {
        "id": None,
        "related_resources": (
            "list",
            {
                0: {
                    "id": None,
                    "related_resources": (
                        "list",
                        {
                            0: {
                                "resource_type": None,
                                "name": None,
                                "patient": None,
                                "practitioner": None,
                            },
                            1: {
                                "resource_type": None,
                                "name": None,
                                "patient": None,
                                "practitioner": None,
                            },
                        },
                    ),
                    "memo": {
                        "description": None,
                        "account": {
                            "resource_type": None,
                            "name": None,
                            "patient": None,
                            "practitioner": None,
                        },
                    },
                    "resource_type": None,
                    "name": None,
                    "patient": {
                        "id": None,
                        "related_resources": ("list", {}),
                        "memo": {"description": None, "account": None},
                        "resource_type": None,
                        "name": None,
                        "scripts": (
                            "list",
                            {
                                0: {"resource_type": None, "name": None},
                                1: {"resource_type": None, "name": None},
                            },
                        ),
                    },
                    "practitioner": {
                        "id": None,
                        "related_resources": ("list", {}),
                        "memo": {"description": None, "account": None},
                        "resource_type": None,
                        "name": None,
                    },
                },
                1: {
                    "id": None,
                    "related_resources": (
                        "list",
                        {
                            0: {
                                "resource_type": None,
                                "name": None,
                                "patient": None,
                                "practitioner": None,
                            },
                            1: {
                                "resource_type": None,
                                "name": None,
                                "patient": None,
                                "practitioner": None,
                            },
                        },
                    ),
                    "memo": {
                        "description": None,
                        "account": {
                            "resource_type": None,
                            "name": None,
                            "patient": None,
                            "practitioner": None,
                        },
                    },
                    "resource_type": None,
                    "name": None,
                    "patient": {
                        "id": None,
                        "related_resources": ("list", {}),
                        "memo": {"description": None, "account": None},
                        "resource_type": None,
                        "name": None,
                        "scripts": (
                            "list",
                            {
                                0: {"resource_type": None, "name": None},
                                1: {"resource_type": None, "name": None},
                            },
                        ),
                    },
                    "practitioner": {
                        "id": None,
                        "related_resources": ("list", {}),
                        "memo": {"description": None, "account": None},
                        "resource_type": None,
                        "name": None,
                    },
                },
            },
        ),
        "memo": {
            "description": None,
            "account": {
                "id": None,
                "related_resources": (
                    "list",
                    {
                        0: {
                            "resource_type": None,
                            "name": None,
                            "patient": None,
                            "practitioner": None,
                        },
                        1: {
                            "resource_type": None,
                            "name": None,
                            "patient": None,
                            "practitioner": None,
                        },
                    },
                ),
                "memo": {
                    "description": None,
                    "account": {
                        "resource_type": None,
                        "name": None,
                        "patient": None,
                        "practitioner": None,
                    },
                },
                "resource_type": None,
                "name": None,
                "patient": {
                    "id": None,
                    "related_resources": ("list", {}),
                    "memo": {"description": None, "account": None},
                    "resource_type": None,
                    "name": None,
                    "scripts": (
                        "list",
                        {
                            0: {"resource_type": None, "name": None},
                            1: {"resource_type": None, "name": None},
                        },
                    ),
                },
                "practitioner": {
                    "id": None,
                    "related_resources": ("list", {}),
                    "memo": {"description": None, "account": None},
                    "resource_type": None,
                    "name": None,
                },
            },
        },
        "resource_type": None,
        "name": None,
        "patient": {
            "id": None,
            "related_resources": (
                "list",
                {
                    0: {
                        "id": None,
                        "related_resources": ("list", {}),
                        "memo": {"description": None, "account": None},
                        "resource_type": None,
                        "name": None,
                        "patient": {
                            "resource_type": None,
                            "name": None,
                            "scripts": ("list", {}),
                        },
                        "practitioner": {"resource_type": None, "name": None},
                    },
                    1: {
                        "id": None,
                        "related_resources": ("list", {}),
                        "memo": {"description": None, "account": None},
                        "resource_type": None,
                        "name": None,
                        "patient": {
                            "resource_type": None,
                            "name": None,
                            "scripts": ("list", {}),
                        },
                        "practitioner": {"resource_type": None, "name": None},
                    },
                },
            ),
            "memo": {
                "description": None,
                "account": {
                    "id": None,
                    "related_resources": ("list", {}),
                    "memo": {"description": None, "account": None},
                    "resource_type": None,
                    "name": None,
                    "patient": {
                        "resource_type": None,
                        "name": None,
                        "scripts": ("list", {}),
                    },
                    "practitioner": {"resource_type": None, "name": None},
                },
            },
            "resource_type": None,
            "name": None,
            "scripts": (
                "list",
                {
                    0: {
                        "id": None,
                        "related_resources": (
                            "list",
                            {
                                0: {
                                    "resource_type": None,
                                    "name": None,
                                    "patient": None,
                                    "practitioner": None,
                                },
                                1: {
                                    "resource_type": None,
                                    "name": None,
                                    "patient": None,
                                    "practitioner": None,
                                },
                            },
                        ),
                        "memo": {
                            "description": None,
                            "account": {
                                "resource_type": None,
                                "name": None,
                                "patient": None,
                                "practitioner": None,
                            },
                        },
                        "resource_type": None,
                        "name": None,
                    },
                    1: {
                        "id": None,
                        "related_resources": (
                            "list",
                            {
                                0: {
                                    "resource_type": None,
                                    "name": None,
                                    "patient": None,
                                    "practitioner": None,
                                },
                                1: {
                                    "resource_type": None,
                                    "name": None,
                                    "patient": None,
                                    "practitioner": None,
                                },
                            },
                        ),
                        "memo": {
                            "description": None,
                            "account": {
                                "resource_type": None,
                                "name": None,
                                "patient": None,
                                "practitioner": None,
                            },
                        },
                        "resource_type": None,
                        "name": None,
                    },
                },
            ),
        },
        "practitioner": {
            "id": None,
            "related_resources": (
                "list",
                {
                    0: {
                        "id": None,
                        "related_resources": ("list", {}),
                        "memo": {"description": None, "account": None},
                        "resource_type": None,
                        "name": None,
                        "patient": {
                            "resource_type": None,
                            "name": None,
                            "scripts": ("list", {}),
                        },
                        "practitioner": {"resource_type": None, "name": None},
                    },
                    1: {
                        "id": None,
                        "related_resources": ("list", {}),
                        "memo": {"description": None, "account": None},
                        "resource_type": None,
                        "name": None,
                        "patient": {
                            "resource_type": None,
                            "name": None,
                            "scripts": ("list", {}),
                        },
                        "practitioner": {"resource_type": None, "name": None},
                    },
                },
            ),
            "memo": {
                "description": None,
                "account": {
                    "id": None,
                    "related_resources": ("list", {}),
                    "memo": {"description": None, "account": None},
                    "resource_type": None,
                    "name": None,
                    "patient": {
                        "resource_type": None,
                        "name": None,
                        "scripts": ("list", {}),
                    },
                    "practitioner": {"resource_type": None, "name": None},
                },
            },
            "resource_type": None,
            "name": None,
        },
    }
    response = client.get_account(account_id="account_id")
    validate_response(response, expected_response, expected_types)

    async_response = await async_client.get_account(account_id="account_id")
    validate_response(async_response, expected_response, expected_types)
