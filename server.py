import socket
from _thread import start_new_thread
import sys

server = '192.168.0.79'
port = 5555

socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

try:
    socket.bind((server, port))

except socket.error as msg:
    str(msg)

# This defines how many clients can connect to the server
socket.listen(2)
print('Waiting for a connection')


def threaded_client(conn):
    conn.send(str.encode('Hello, client!'))
    reply = str()
    while True:
        try:
            # Bytes that are sent through the thread to here to the server. Thread diameter so to say
            data = conn.recv(2048)
            reply = data.decode('utf-8')

            if not data:
                print('Disconnected')
                break
            else:
                print('Received: ', reply)
                print('Sending reply: ', reply)
            # Encodes string back to binary when server sends back response
            conn.sendall(str.encode(reply))
        except:
            break

    print('Lost connection')
    conn.close()

while True:
    conn, addr = socket.accept()
    print('Connected by', addr)
    # This starts a thread and keeps connection. While this is happening, another thread can start besides that.
    start_new_thread(threaded_client, (conn,))
