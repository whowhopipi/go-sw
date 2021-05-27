
static __thread int tls;

void
setTLS(int v)
{
	tls = v;
}


extern int Ttt();
 int
getTLS()
{
  Ttt(4321);
	return tls;
}

