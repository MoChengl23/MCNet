using System;
using System.Net;
using System.Net.Sockets;

namespace IOCPNet
{
    public class IOCPClient
    {
        Socket socket;
        SocketAsyncEventArgs saea;
        public IOCPClient(){
            saea = new SocketAsyncEventArgs();
        }
        public void StartAsClient(string ip, int port){
            IPEndPoint port = new IPEndPoint(IPAddress.Parse(ip), port);
            socket = new Socket(port.AddressFamily, SocketType.Stream, ProtocolType.Tcp);
            Console.WriteLine("Client Start.");
            StartConnect();
        }
        void StartConnect(){
            //
            bool suspend = socket.ConnectAsync(saea);
            if(!suspend){
                ProcessConnect();
            }

        }

        void ProcessConnect(){

        }
        


    }
}