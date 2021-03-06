import lzma
import sys

import bson  # from PyMongo


with open("commits.bin", "wb") as commits:
    with lzma.open("repos.txt.xz", "wb") as repos:
        with lzma.open("messages.txt.xz", "wb") as messages:
            while True:
                try:
                    for obj in bson.decode_file_iter(sys.stdin.buffer, codec_options=bson.CodecOptions(
                            unicode_decode_error_handler="ignore")):
                        try:
                            obj["sha"]
                            obj["commit"]["url"]
                            obj["commit"]["message"]
                        except KeyError:
                            continue
                        commits.write(bytes.fromhex(obj["sha"]))
                        repos.write(obj["commit"]["url"][29:-53].encode())
                        repos.write(b"\0")
                        messages.write(obj["commit"]["message"].encode(errors="ignore"))
                        messages.write(b"\0")
                    break
                except bson.errors.InvalidBSON:
                    continue
