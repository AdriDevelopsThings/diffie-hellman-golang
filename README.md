# diffie-hellman-golang

I tried to implement the diffie hellman key exchange in golang with the use of sockets to understand how this protocol works and to learn golang. So its a development project: obviously don't use it in production.

## Installation

Run ``go get`` to download the needed dependencies.

## Run

Run this with ``go run .``
You need to set the following command arguments:

| Name            | Help                                        |
| --------------- | ------------------------------------------- |
| -m or --mode    | mode of diffie hellman `server` or `client` |
| -a or --address | the listen address or address of the server |

## How it works

After the connection between the server and client is established the client will generate the public numbers p and g. p is a primnumber and g is lower than p. The client also generates/calculates his secret `s` and public key. The secret is smaller than p and the public key is g ^ s mod p. After the server received the public information he generates/calculates his secret and public key as well. He sends his public key to the client and the client sends his back. Both sides calculate the shared secret: The public key of the other partner ^ the own secret `s` mod p. The clients generated a shared secret asymmetricaly.