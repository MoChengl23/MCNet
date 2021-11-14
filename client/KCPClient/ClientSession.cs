
using KCPNET;
using Pb;
using System;

namespace KCPClient
{
    public class ClientSession : Session 
    {
        
        protected override void OnConnected() {
        }

        protected override void OnDisConnected() {
        }

        protected override void OnReceiveMessage(PbMessage pbMessage) {
            Console.WriteLine(pbMessage);
        }

        protected override void OnUpdate(DateTime now) {
        }
    
    }
}