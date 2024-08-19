import socket
from _thread import start_new_thread
import sys

server = '10.252.11.186'
port = 5555

socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

try:
    socket.bind((server, port))

except socket.error as msg:
    str(msg)

# This defines how many clients can connect to the server
socket.listen(2)
print(f'Waiting for a connection. Server listens to port {port}')

def read_position(str):
    str = str.split(',')
    return int(str[0]), int(str[1])

def make_position(tup):
    return str(tup[0]) + ',' + str(tup[1])

position = [(0,0), (100,100)]

def threaded_client(conn, current_player):
    conn.send(str.encode(make_position(position[current_player])))
    reply = str()
    while True:
        try:
            # Bytes that are sent through the thread to here to the server. Thread diameter so to say
            data = read_position(conn.recv(2048).decode())
            position[current_player] = data


            if not data:
                print('Disconnected')
                break
            else:
                if current_player == 1:
                    reply = position[0]
                else:
                    reply = position[1]
                print('Received: ', data)
                print('Sending reply: ', reply)
            # Encodes string back to binary when server sends back response
            conn.sendall(str.encode(make_position(reply)))
        except:
            break

    print('Lost connection')
    conn.close()

current_player = 0

while True:
    conn, addr = socket.accept()
    print('Connected by', addr)
    # This starts a thread and keeps connection. While this is happening, another thread can start besides that.
    start_new_thread(threaded_client, (conn, current_player))
    current_player += 1
