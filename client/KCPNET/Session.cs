  
using System;
using System.Buffers;
using System.Net.Sockets.Kcp;
using System.Text;
using System.Threading;
using System.Threading.Tasks;
using System.Net;
using Google.Protobuf;
using System.Net.Sockets;
using Pb;





namespace KCPNET
{
    public abstract class Session<T>
    where T: Msg , new(){

        private enum State{
            DisConnected,
            Connected,
        }
        private State state = State.DisConnected;
        public bool IsConnected {get{return state == State.Connected;}}
        protected uint m_sid;
       
        public Kcp m_kcp;
        public UdpClient udp;
        Handle m_handle;
        public void InitSession(uint sid,UdpClient _udp, IPEndPoint remotePoint)
        
        {
            udp = _udp;
            state = State.Connected;
            m_sid = sid;
            m_handle = new Handle();
            m_kcp = new Kcp(sid, m_handle);
            m_kcp.NoDelay(1, 10, 2, 1);
            m_kcp.WndSize(64, 64);
            m_kcp.SetMtu(512);
       
            m_handle.Out  += buffer =>
            {
                byte[] bytes = buffer.ToArray();
                udp.SendAsync(bytes, bytes.Length, remotePoint);       
            };
            m_handle.Recv += buffer =>
            {
                

                
                Console.WriteLine("触发handle.recv");
                Console.WriteLine(Encoding.UTF8.GetString(buffer));
                // ReadFromPb(buffer);
                // var a = SendByPb();
                // udp.SendAsync(a, a.Length,new IPEndPoint(IPAddress.Parse("127.0.0.1"), 7777));
            };
            Task.Run(Update);
        }
        /// <summary>
        /// 将二进制数据转化成protobuf格式
        /// 且一定要用try
        /// </summary>
        /// <param name="bytes"></param>
        // void ReadFromPb(byte[] bytes)
        // {
        //     try{
        //         var person = new Student{};
        //         var me = Student.Parser.ParseFrom(bytes);
        //         Console.WriteLine(me);
        //     }
        //     catch(Exception e)
        //     {
                
        //     }
            
           
        // }


        /// <summary>
        /// 将protobuf结构体的形式转化成bytes
        /// </summary>
        /// <returns></returns>
        byte[] PbToByte(T msg)
        {
            byte[] bytes = new byte[128];
            try{
                var pbMessage = new PbMessage{
                    Name = msg.name,
                };
                
                CodedOutputStream output  = new CodedOutputStream(bytes);
                pbMessage.WriteTo(output);
          
            }
            catch(Exception e)
            {
                
            }
            return bytes;
            
        }


        public void ReceiveData(byte[] bytes){
            m_kcp.Input(bytes.AsSpan());

        }
        /// <summary>
        /// 规定，如果发送长度为1的byte，表示玩家进入对局，长度为2表示玩家退出对局
        /// </summary>
        /// <param name="msg"></param>
        public void SendMessage(T msg){
            
            if(IsConnected)
                m_kcp.Send(PbToByte(msg));
                //  m_kcp.Send(new byte[]{1});
        }


        // async public void ServerReceive()
        // {
        //     //我是服务器,我接受udp的端口是6666,我发射的端口随机
            
            
        //     UdpReceiveResult result;
        //     while(true)
        //     {
              
        //         result = await udp.ReceiveAsync();
        //          byte[] bytes = result.Buffer.AsSpan().ToArray();
        //         //  for(int i=0;i<bytes.Length;i++)
        //         // {
        //         //    Console.Write(bytes[i]+ " ");
        //         // }
        //         // Console.WriteLine(BitConverter.ToUInt32(result.Buffer, 0) );
                
        //         // Console.WriteLine(result.RemoteEndPoint);
        //         if(!isConnected)
        //         {
        //             InitSession(3, result.RemoteEndPoint);
        //             isConnected = true;
        //         }
        //         else
        //         {
                    
        //             m_kcp.Input(result.Buffer.AsSpan());
                 
                 
              

        //         }
                
        //     }
        // }
        
    async public void Update()
        {
            try
            {
                while (true)
                {

                        m_kcp.Update(DateTime.UtcNow);
                        int len;
                        do
                        {
                           
                            var (buffer, avalidSzie) = m_kcp.TryRecv();
                            len = avalidSzie;
                            if (buffer != null)
                            {
                              
                                
                                var temp = new byte[len];
                                buffer.Memory.Span.Slice(0, len).CopyTo(temp);
                                m_handle.Receive(temp);
                            }
                        } while (len > 0);
                    
                        await Task.Delay(10);
                    
                }
            }
            catch (Exception e)
            {

            }
        }

    

    }
}