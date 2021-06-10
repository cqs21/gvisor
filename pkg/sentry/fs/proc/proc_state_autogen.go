// automatically generated by stateify.

package proc

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (i *execArgInode) StateTypeName() string {
	return "pkg/sentry/fs/proc.execArgInode"
}

func (i *execArgInode) StateFields() []string {
	return []string{
		"SimpleFileInode",
		"arg",
		"t",
	}
}

func (i *execArgInode) beforeSave() {}

// +checklocksignore
func (i *execArgInode) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.SimpleFileInode)
	stateSinkObject.Save(1, &i.arg)
	stateSinkObject.Save(2, &i.t)
}

func (i *execArgInode) afterLoad() {}

// +checklocksignore
func (i *execArgInode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.SimpleFileInode)
	stateSourceObject.Load(1, &i.arg)
	stateSourceObject.Load(2, &i.t)
}

func (f *execArgFile) StateTypeName() string {
	return "pkg/sentry/fs/proc.execArgFile"
}

func (f *execArgFile) StateFields() []string {
	return []string{
		"arg",
		"t",
	}
}

func (f *execArgFile) beforeSave() {}

// +checklocksignore
func (f *execArgFile) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.arg)
	stateSinkObject.Save(1, &f.t)
}

func (f *execArgFile) afterLoad() {}

// +checklocksignore
func (f *execArgFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.arg)
	stateSourceObject.Load(1, &f.t)
}

func (f *fdDir) StateTypeName() string {
	return "pkg/sentry/fs/proc.fdDir"
}

func (f *fdDir) StateFields() []string {
	return []string{
		"Dir",
		"t",
	}
}

func (f *fdDir) beforeSave() {}

// +checklocksignore
func (f *fdDir) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.Dir)
	stateSinkObject.Save(1, &f.t)
}

func (f *fdDir) afterLoad() {}

// +checklocksignore
func (f *fdDir) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.Dir)
	stateSourceObject.Load(1, &f.t)
}

func (f *fdDirFile) StateTypeName() string {
	return "pkg/sentry/fs/proc.fdDirFile"
}

func (f *fdDirFile) StateFields() []string {
	return []string{
		"isInfoFile",
		"t",
	}
}

func (f *fdDirFile) beforeSave() {}

// +checklocksignore
func (f *fdDirFile) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.isInfoFile)
	stateSinkObject.Save(1, &f.t)
}

func (f *fdDirFile) afterLoad() {}

// +checklocksignore
func (f *fdDirFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.isInfoFile)
	stateSourceObject.Load(1, &f.t)
}

func (fdid *fdInfoDir) StateTypeName() string {
	return "pkg/sentry/fs/proc.fdInfoDir"
}

func (fdid *fdInfoDir) StateFields() []string {
	return []string{
		"Dir",
		"t",
	}
}

func (fdid *fdInfoDir) beforeSave() {}

// +checklocksignore
func (fdid *fdInfoDir) StateSave(stateSinkObject state.Sink) {
	fdid.beforeSave()
	stateSinkObject.Save(0, &fdid.Dir)
	stateSinkObject.Save(1, &fdid.t)
}

func (fdid *fdInfoDir) afterLoad() {}

// +checklocksignore
func (fdid *fdInfoDir) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fdid.Dir)
	stateSourceObject.Load(1, &fdid.t)
}

func (f *filesystemsData) StateTypeName() string {
	return "pkg/sentry/fs/proc.filesystemsData"
}

func (f *filesystemsData) StateFields() []string {
	return []string{}
}

func (f *filesystemsData) beforeSave() {}

// +checklocksignore
func (f *filesystemsData) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
}

func (f *filesystemsData) afterLoad() {}

// +checklocksignore
func (f *filesystemsData) StateLoad(stateSourceObject state.Source) {
}

func (f *filesystem) StateTypeName() string {
	return "pkg/sentry/fs/proc.filesystem"
}

func (f *filesystem) StateFields() []string {
	return []string{}
}

func (f *filesystem) beforeSave() {}

// +checklocksignore
func (f *filesystem) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
}

func (f *filesystem) afterLoad() {}

// +checklocksignore
func (f *filesystem) StateLoad(stateSourceObject state.Source) {
}

func (i *taskOwnedInodeOps) StateTypeName() string {
	return "pkg/sentry/fs/proc.taskOwnedInodeOps"
}

func (i *taskOwnedInodeOps) StateFields() []string {
	return []string{
		"InodeOperations",
		"t",
	}
}

func (i *taskOwnedInodeOps) beforeSave() {}

// +checklocksignore
func (i *taskOwnedInodeOps) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.InodeOperations)
	stateSinkObject.Save(1, &i.t)
}

func (i *taskOwnedInodeOps) afterLoad() {}

// +checklocksignore
func (i *taskOwnedInodeOps) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.InodeOperations)
	stateSourceObject.Load(1, &i.t)
}

func (s *staticFileInodeOps) StateTypeName() string {
	return "pkg/sentry/fs/proc.staticFileInodeOps"
}

func (s *staticFileInodeOps) StateFields() []string {
	return []string{
		"InodeSimpleAttributes",
		"InodeStaticFileGetter",
	}
}

func (s *staticFileInodeOps) beforeSave() {}

// +checklocksignore
func (s *staticFileInodeOps) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.InodeSimpleAttributes)
	stateSinkObject.Save(1, &s.InodeStaticFileGetter)
}

func (s *staticFileInodeOps) afterLoad() {}

// +checklocksignore
func (s *staticFileInodeOps) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.InodeSimpleAttributes)
	stateSourceObject.Load(1, &s.InodeStaticFileGetter)
}

func (d *loadavgData) StateTypeName() string {
	return "pkg/sentry/fs/proc.loadavgData"
}

func (d *loadavgData) StateFields() []string {
	return []string{}
}

func (d *loadavgData) beforeSave() {}

// +checklocksignore
func (d *loadavgData) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
}

func (d *loadavgData) afterLoad() {}

// +checklocksignore
func (d *loadavgData) StateLoad(stateSourceObject state.Source) {
}

func (d *meminfoData) StateTypeName() string {
	return "pkg/sentry/fs/proc.meminfoData"
}

func (d *meminfoData) StateFields() []string {
	return []string{
		"k",
	}
}

func (d *meminfoData) beforeSave() {}

// +checklocksignore
func (d *meminfoData) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.k)
}

func (d *meminfoData) afterLoad() {}

// +checklocksignore
func (d *meminfoData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.k)
}

func (mif *mountInfoFile) StateTypeName() string {
	return "pkg/sentry/fs/proc.mountInfoFile"
}

func (mif *mountInfoFile) StateFields() []string {
	return []string{
		"t",
	}
}

func (mif *mountInfoFile) beforeSave() {}

// +checklocksignore
func (mif *mountInfoFile) StateSave(stateSinkObject state.Sink) {
	mif.beforeSave()
	stateSinkObject.Save(0, &mif.t)
}

func (mif *mountInfoFile) afterLoad() {}

// +checklocksignore
func (mif *mountInfoFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &mif.t)
}

func (mf *mountsFile) StateTypeName() string {
	return "pkg/sentry/fs/proc.mountsFile"
}

func (mf *mountsFile) StateFields() []string {
	return []string{
		"t",
	}
}

func (mf *mountsFile) beforeSave() {}

// +checklocksignore
func (mf *mountsFile) StateSave(stateSinkObject state.Sink) {
	mf.beforeSave()
	stateSinkObject.Save(0, &mf.t)
}

func (mf *mountsFile) afterLoad() {}

// +checklocksignore
func (mf *mountsFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &mf.t)
}

func (n *ifinet6) StateTypeName() string {
	return "pkg/sentry/fs/proc.ifinet6"
}

func (n *ifinet6) StateFields() []string {
	return []string{
		"s",
	}
}

func (n *ifinet6) beforeSave() {}

// +checklocksignore
func (n *ifinet6) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.s)
}

func (n *ifinet6) afterLoad() {}

// +checklocksignore
func (n *ifinet6) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.s)
}

func (n *netDev) StateTypeName() string {
	return "pkg/sentry/fs/proc.netDev"
}

func (n *netDev) StateFields() []string {
	return []string{
		"s",
	}
}

func (n *netDev) beforeSave() {}

// +checklocksignore
func (n *netDev) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.s)
}

func (n *netDev) afterLoad() {}

// +checklocksignore
func (n *netDev) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.s)
}

func (n *netSnmp) StateTypeName() string {
	return "pkg/sentry/fs/proc.netSnmp"
}

func (n *netSnmp) StateFields() []string {
	return []string{
		"s",
	}
}

func (n *netSnmp) beforeSave() {}

// +checklocksignore
func (n *netSnmp) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.s)
}

func (n *netSnmp) afterLoad() {}

// +checklocksignore
func (n *netSnmp) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.s)
}

func (n *netRoute) StateTypeName() string {
	return "pkg/sentry/fs/proc.netRoute"
}

func (n *netRoute) StateFields() []string {
	return []string{
		"s",
	}
}

func (n *netRoute) beforeSave() {}

// +checklocksignore
func (n *netRoute) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.s)
}

func (n *netRoute) afterLoad() {}

// +checklocksignore
func (n *netRoute) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.s)
}

func (n *netUnix) StateTypeName() string {
	return "pkg/sentry/fs/proc.netUnix"
}

func (n *netUnix) StateFields() []string {
	return []string{
		"k",
	}
}

func (n *netUnix) beforeSave() {}

// +checklocksignore
func (n *netUnix) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.k)
}

func (n *netUnix) afterLoad() {}

// +checklocksignore
func (n *netUnix) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.k)
}

func (n *netTCP) StateTypeName() string {
	return "pkg/sentry/fs/proc.netTCP"
}

func (n *netTCP) StateFields() []string {
	return []string{
		"k",
	}
}

func (n *netTCP) beforeSave() {}

// +checklocksignore
func (n *netTCP) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.k)
}

func (n *netTCP) afterLoad() {}

// +checklocksignore
func (n *netTCP) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.k)
}

func (n *netTCP6) StateTypeName() string {
	return "pkg/sentry/fs/proc.netTCP6"
}

func (n *netTCP6) StateFields() []string {
	return []string{
		"k",
	}
}

func (n *netTCP6) beforeSave() {}

// +checklocksignore
func (n *netTCP6) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.k)
}

func (n *netTCP6) afterLoad() {}

// +checklocksignore
func (n *netTCP6) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.k)
}

func (n *netUDP) StateTypeName() string {
	return "pkg/sentry/fs/proc.netUDP"
}

func (n *netUDP) StateFields() []string {
	return []string{
		"k",
	}
}

func (n *netUDP) beforeSave() {}

// +checklocksignore
func (n *netUDP) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.k)
}

func (n *netUDP) afterLoad() {}

// +checklocksignore
func (n *netUDP) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.k)
}

func (p *proc) StateTypeName() string {
	return "pkg/sentry/fs/proc.proc"
}

func (p *proc) StateFields() []string {
	return []string{
		"Dir",
		"k",
		"pidns",
		"cgroupControllers",
	}
}

func (p *proc) beforeSave() {}

// +checklocksignore
func (p *proc) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	stateSinkObject.Save(0, &p.Dir)
	stateSinkObject.Save(1, &p.k)
	stateSinkObject.Save(2, &p.pidns)
	stateSinkObject.Save(3, &p.cgroupControllers)
}

func (p *proc) afterLoad() {}

// +checklocksignore
func (p *proc) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.Dir)
	stateSourceObject.Load(1, &p.k)
	stateSourceObject.Load(2, &p.pidns)
	stateSourceObject.Load(3, &p.cgroupControllers)
}

func (s *self) StateTypeName() string {
	return "pkg/sentry/fs/proc.self"
}

func (s *self) StateFields() []string {
	return []string{
		"Symlink",
		"pidns",
	}
}

func (s *self) beforeSave() {}

// +checklocksignore
func (s *self) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.Symlink)
	stateSinkObject.Save(1, &s.pidns)
}

func (s *self) afterLoad() {}

// +checklocksignore
func (s *self) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.Symlink)
	stateSourceObject.Load(1, &s.pidns)
}

func (s *threadSelf) StateTypeName() string {
	return "pkg/sentry/fs/proc.threadSelf"
}

func (s *threadSelf) StateFields() []string {
	return []string{
		"Symlink",
		"pidns",
	}
}

func (s *threadSelf) beforeSave() {}

// +checklocksignore
func (s *threadSelf) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.Symlink)
	stateSinkObject.Save(1, &s.pidns)
}

func (s *threadSelf) afterLoad() {}

// +checklocksignore
func (s *threadSelf) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.Symlink)
	stateSourceObject.Load(1, &s.pidns)
}

func (rpf *rootProcFile) StateTypeName() string {
	return "pkg/sentry/fs/proc.rootProcFile"
}

func (rpf *rootProcFile) StateFields() []string {
	return []string{
		"iops",
	}
}

func (rpf *rootProcFile) beforeSave() {}

// +checklocksignore
func (rpf *rootProcFile) StateSave(stateSinkObject state.Sink) {
	rpf.beforeSave()
	stateSinkObject.Save(0, &rpf.iops)
}

func (rpf *rootProcFile) afterLoad() {}

// +checklocksignore
func (rpf *rootProcFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &rpf.iops)
}

func (s *statData) StateTypeName() string {
	return "pkg/sentry/fs/proc.statData"
}

func (s *statData) StateFields() []string {
	return []string{
		"k",
	}
}

func (s *statData) beforeSave() {}

// +checklocksignore
func (s *statData) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.k)
}

func (s *statData) afterLoad() {}

// +checklocksignore
func (s *statData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.k)
}

func (d *mmapMinAddrData) StateTypeName() string {
	return "pkg/sentry/fs/proc.mmapMinAddrData"
}

func (d *mmapMinAddrData) StateFields() []string {
	return []string{
		"k",
	}
}

func (d *mmapMinAddrData) beforeSave() {}

// +checklocksignore
func (d *mmapMinAddrData) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.k)
}

func (d *mmapMinAddrData) afterLoad() {}

// +checklocksignore
func (d *mmapMinAddrData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.k)
}

func (o *overcommitMemory) StateTypeName() string {
	return "pkg/sentry/fs/proc.overcommitMemory"
}

func (o *overcommitMemory) StateFields() []string {
	return []string{}
}

func (o *overcommitMemory) beforeSave() {}

// +checklocksignore
func (o *overcommitMemory) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
}

func (o *overcommitMemory) afterLoad() {}

// +checklocksignore
func (o *overcommitMemory) StateLoad(stateSourceObject state.Source) {
}

func (m *maxMapCount) StateTypeName() string {
	return "pkg/sentry/fs/proc.maxMapCount"
}

func (m *maxMapCount) StateFields() []string {
	return []string{}
}

func (m *maxMapCount) beforeSave() {}

// +checklocksignore
func (m *maxMapCount) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
}

func (m *maxMapCount) afterLoad() {}

// +checklocksignore
func (m *maxMapCount) StateLoad(stateSourceObject state.Source) {
}

func (h *hostname) StateTypeName() string {
	return "pkg/sentry/fs/proc.hostname"
}

func (h *hostname) StateFields() []string {
	return []string{
		"SimpleFileInode",
	}
}

func (h *hostname) beforeSave() {}

// +checklocksignore
func (h *hostname) StateSave(stateSinkObject state.Sink) {
	h.beforeSave()
	stateSinkObject.Save(0, &h.SimpleFileInode)
}

func (h *hostname) afterLoad() {}

// +checklocksignore
func (h *hostname) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &h.SimpleFileInode)
}

func (hf *hostnameFile) StateTypeName() string {
	return "pkg/sentry/fs/proc.hostnameFile"
}

func (hf *hostnameFile) StateFields() []string {
	return []string{}
}

func (hf *hostnameFile) beforeSave() {}

// +checklocksignore
func (hf *hostnameFile) StateSave(stateSinkObject state.Sink) {
	hf.beforeSave()
}

func (hf *hostnameFile) afterLoad() {}

// +checklocksignore
func (hf *hostnameFile) StateLoad(stateSourceObject state.Source) {
}

func (t *tcpMemInode) StateTypeName() string {
	return "pkg/sentry/fs/proc.tcpMemInode"
}

func (t *tcpMemInode) StateFields() []string {
	return []string{
		"SimpleFileInode",
		"dir",
		"s",
		"size",
	}
}

// +checklocksignore
func (t *tcpMemInode) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.SimpleFileInode)
	stateSinkObject.Save(1, &t.dir)
	stateSinkObject.Save(2, &t.s)
	stateSinkObject.Save(3, &t.size)
}

// +checklocksignore
func (t *tcpMemInode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.SimpleFileInode)
	stateSourceObject.Load(1, &t.dir)
	stateSourceObject.LoadWait(2, &t.s)
	stateSourceObject.Load(3, &t.size)
	stateSourceObject.AfterLoad(t.afterLoad)
}

func (f *tcpMemFile) StateTypeName() string {
	return "pkg/sentry/fs/proc.tcpMemFile"
}

func (f *tcpMemFile) StateFields() []string {
	return []string{
		"tcpMemInode",
	}
}

func (f *tcpMemFile) beforeSave() {}

// +checklocksignore
func (f *tcpMemFile) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.tcpMemInode)
}

func (f *tcpMemFile) afterLoad() {}

// +checklocksignore
func (f *tcpMemFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.tcpMemInode)
}

func (s *tcpSack) StateTypeName() string {
	return "pkg/sentry/fs/proc.tcpSack"
}

func (s *tcpSack) StateFields() []string {
	return []string{
		"SimpleFileInode",
		"stack",
		"enabled",
	}
}

func (s *tcpSack) beforeSave() {}

// +checklocksignore
func (s *tcpSack) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.SimpleFileInode)
	stateSinkObject.Save(1, &s.stack)
	stateSinkObject.Save(2, &s.enabled)
}

// +checklocksignore
func (s *tcpSack) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.SimpleFileInode)
	stateSourceObject.LoadWait(1, &s.stack)
	stateSourceObject.Load(2, &s.enabled)
	stateSourceObject.AfterLoad(s.afterLoad)
}

func (f *tcpSackFile) StateTypeName() string {
	return "pkg/sentry/fs/proc.tcpSackFile"
}

func (f *tcpSackFile) StateFields() []string {
	return []string{
		"tcpSack",
		"stack",
	}
}

func (f *tcpSackFile) beforeSave() {}

// +checklocksignore
func (f *tcpSackFile) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.tcpSack)
	stateSinkObject.Save(1, &f.stack)
}

func (f *tcpSackFile) afterLoad() {}

// +checklocksignore
func (f *tcpSackFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.tcpSack)
	stateSourceObject.LoadWait(1, &f.stack)
}

func (r *tcpRecovery) StateTypeName() string {
	return "pkg/sentry/fs/proc.tcpRecovery"
}

func (r *tcpRecovery) StateFields() []string {
	return []string{
		"SimpleFileInode",
		"stack",
		"recovery",
	}
}

func (r *tcpRecovery) beforeSave() {}

// +checklocksignore
func (r *tcpRecovery) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.SimpleFileInode)
	stateSinkObject.Save(1, &r.stack)
	stateSinkObject.Save(2, &r.recovery)
}

func (r *tcpRecovery) afterLoad() {}

// +checklocksignore
func (r *tcpRecovery) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.SimpleFileInode)
	stateSourceObject.LoadWait(1, &r.stack)
	stateSourceObject.Load(2, &r.recovery)
}

func (f *tcpRecoveryFile) StateTypeName() string {
	return "pkg/sentry/fs/proc.tcpRecoveryFile"
}

func (f *tcpRecoveryFile) StateFields() []string {
	return []string{
		"tcpRecovery",
		"stack",
	}
}

func (f *tcpRecoveryFile) beforeSave() {}

// +checklocksignore
func (f *tcpRecoveryFile) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.tcpRecovery)
	stateSinkObject.Save(1, &f.stack)
}

func (f *tcpRecoveryFile) afterLoad() {}

// +checklocksignore
func (f *tcpRecoveryFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.tcpRecovery)
	stateSourceObject.LoadWait(1, &f.stack)
}

func (ipf *ipForwarding) StateTypeName() string {
	return "pkg/sentry/fs/proc.ipForwarding"
}

func (ipf *ipForwarding) StateFields() []string {
	return []string{
		"SimpleFileInode",
		"stack",
		"enabled",
	}
}

func (ipf *ipForwarding) beforeSave() {}

// +checklocksignore
func (ipf *ipForwarding) StateSave(stateSinkObject state.Sink) {
	ipf.beforeSave()
	stateSinkObject.Save(0, &ipf.SimpleFileInode)
	stateSinkObject.Save(1, &ipf.stack)
	stateSinkObject.Save(2, &ipf.enabled)
}

// +checklocksignore
func (ipf *ipForwarding) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &ipf.SimpleFileInode)
	stateSourceObject.LoadWait(1, &ipf.stack)
	stateSourceObject.Load(2, &ipf.enabled)
	stateSourceObject.AfterLoad(ipf.afterLoad)
}

func (f *ipForwardingFile) StateTypeName() string {
	return "pkg/sentry/fs/proc.ipForwardingFile"
}

func (f *ipForwardingFile) StateFields() []string {
	return []string{
		"ipf",
		"stack",
	}
}

func (f *ipForwardingFile) beforeSave() {}

// +checklocksignore
func (f *ipForwardingFile) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.ipf)
	stateSinkObject.Save(1, &f.stack)
}

func (f *ipForwardingFile) afterLoad() {}

// +checklocksignore
func (f *ipForwardingFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.ipf)
	stateSourceObject.LoadWait(1, &f.stack)
}

func (in *portRangeInode) StateTypeName() string {
	return "pkg/sentry/fs/proc.portRangeInode"
}

func (in *portRangeInode) StateFields() []string {
	return []string{
		"SimpleFileInode",
		"stack",
		"start",
		"end",
	}
}

func (in *portRangeInode) beforeSave() {}

// +checklocksignore
func (in *portRangeInode) StateSave(stateSinkObject state.Sink) {
	in.beforeSave()
	stateSinkObject.Save(0, &in.SimpleFileInode)
	stateSinkObject.Save(1, &in.stack)
	stateSinkObject.Save(2, &in.start)
	stateSinkObject.Save(3, &in.end)
}

func (in *portRangeInode) afterLoad() {}

// +checklocksignore
func (in *portRangeInode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &in.SimpleFileInode)
	stateSourceObject.LoadWait(1, &in.stack)
	stateSourceObject.Load(2, &in.start)
	stateSourceObject.Load(3, &in.end)
}

func (pf *portRangeFile) StateTypeName() string {
	return "pkg/sentry/fs/proc.portRangeFile"
}

func (pf *portRangeFile) StateFields() []string {
	return []string{
		"inode",
	}
}

func (pf *portRangeFile) beforeSave() {}

// +checklocksignore
func (pf *portRangeFile) StateSave(stateSinkObject state.Sink) {
	pf.beforeSave()
	stateSinkObject.Save(0, &pf.inode)
}

func (pf *portRangeFile) afterLoad() {}

// +checklocksignore
func (pf *portRangeFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &pf.inode)
}

func (t *taskDir) StateTypeName() string {
	return "pkg/sentry/fs/proc.taskDir"
}

func (t *taskDir) StateFields() []string {
	return []string{
		"Dir",
		"t",
	}
}

func (t *taskDir) beforeSave() {}

// +checklocksignore
func (t *taskDir) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.Dir)
	stateSinkObject.Save(1, &t.t)
}

func (t *taskDir) afterLoad() {}

// +checklocksignore
func (t *taskDir) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.Dir)
	stateSourceObject.Load(1, &t.t)
}

func (s *subtasks) StateTypeName() string {
	return "pkg/sentry/fs/proc.subtasks"
}

func (s *subtasks) StateFields() []string {
	return []string{
		"Dir",
		"t",
		"p",
	}
}

func (s *subtasks) beforeSave() {}

// +checklocksignore
func (s *subtasks) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.Dir)
	stateSinkObject.Save(1, &s.t)
	stateSinkObject.Save(2, &s.p)
}

func (s *subtasks) afterLoad() {}

// +checklocksignore
func (s *subtasks) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.Dir)
	stateSourceObject.Load(1, &s.t)
	stateSourceObject.Load(2, &s.p)
}

func (f *subtasksFile) StateTypeName() string {
	return "pkg/sentry/fs/proc.subtasksFile"
}

func (f *subtasksFile) StateFields() []string {
	return []string{
		"t",
		"pidns",
	}
}

func (f *subtasksFile) beforeSave() {}

// +checklocksignore
func (f *subtasksFile) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.t)
	stateSinkObject.Save(1, &f.pidns)
}

func (f *subtasksFile) afterLoad() {}

// +checklocksignore
func (f *subtasksFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.t)
	stateSourceObject.Load(1, &f.pidns)
}

func (e *exe) StateTypeName() string {
	return "pkg/sentry/fs/proc.exe"
}

func (e *exe) StateFields() []string {
	return []string{
		"Symlink",
		"t",
	}
}

func (e *exe) beforeSave() {}

// +checklocksignore
func (e *exe) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.Symlink)
	stateSinkObject.Save(1, &e.t)
}

func (e *exe) afterLoad() {}

// +checklocksignore
func (e *exe) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.Symlink)
	stateSourceObject.Load(1, &e.t)
}

func (e *cwd) StateTypeName() string {
	return "pkg/sentry/fs/proc.cwd"
}

func (e *cwd) StateFields() []string {
	return []string{
		"Symlink",
		"t",
	}
}

func (e *cwd) beforeSave() {}

// +checklocksignore
func (e *cwd) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.Symlink)
	stateSinkObject.Save(1, &e.t)
}

func (e *cwd) afterLoad() {}

// +checklocksignore
func (e *cwd) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.Symlink)
	stateSourceObject.Load(1, &e.t)
}

func (n *namespaceSymlink) StateTypeName() string {
	return "pkg/sentry/fs/proc.namespaceSymlink"
}

func (n *namespaceSymlink) StateFields() []string {
	return []string{
		"Symlink",
		"t",
	}
}

func (n *namespaceSymlink) beforeSave() {}

// +checklocksignore
func (n *namespaceSymlink) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.Symlink)
	stateSinkObject.Save(1, &n.t)
}

func (n *namespaceSymlink) afterLoad() {}

// +checklocksignore
func (n *namespaceSymlink) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.Symlink)
	stateSourceObject.Load(1, &n.t)
}

func (m *memData) StateTypeName() string {
	return "pkg/sentry/fs/proc.memData"
}

func (m *memData) StateFields() []string {
	return []string{
		"SimpleFileInode",
		"t",
	}
}

func (m *memData) beforeSave() {}

// +checklocksignore
func (m *memData) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.SimpleFileInode)
	stateSinkObject.Save(1, &m.t)
}

func (m *memData) afterLoad() {}

// +checklocksignore
func (m *memData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.SimpleFileInode)
	stateSourceObject.Load(1, &m.t)
}

func (m *memDataFile) StateTypeName() string {
	return "pkg/sentry/fs/proc.memDataFile"
}

func (m *memDataFile) StateFields() []string {
	return []string{
		"t",
	}
}

func (m *memDataFile) beforeSave() {}

// +checklocksignore
func (m *memDataFile) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.t)
}

func (m *memDataFile) afterLoad() {}

// +checklocksignore
func (m *memDataFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.t)
}

func (md *mapsData) StateTypeName() string {
	return "pkg/sentry/fs/proc.mapsData"
}

func (md *mapsData) StateFields() []string {
	return []string{
		"t",
	}
}

func (md *mapsData) beforeSave() {}

// +checklocksignore
func (md *mapsData) StateSave(stateSinkObject state.Sink) {
	md.beforeSave()
	stateSinkObject.Save(0, &md.t)
}

func (md *mapsData) afterLoad() {}

// +checklocksignore
func (md *mapsData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &md.t)
}

func (sd *smapsData) StateTypeName() string {
	return "pkg/sentry/fs/proc.smapsData"
}

func (sd *smapsData) StateFields() []string {
	return []string{
		"t",
	}
}

func (sd *smapsData) beforeSave() {}

// +checklocksignore
func (sd *smapsData) StateSave(stateSinkObject state.Sink) {
	sd.beforeSave()
	stateSinkObject.Save(0, &sd.t)
}

func (sd *smapsData) afterLoad() {}

// +checklocksignore
func (sd *smapsData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &sd.t)
}

func (s *taskStatData) StateTypeName() string {
	return "pkg/sentry/fs/proc.taskStatData"
}

func (s *taskStatData) StateFields() []string {
	return []string{
		"t",
		"tgstats",
		"pidns",
	}
}

func (s *taskStatData) beforeSave() {}

// +checklocksignore
func (s *taskStatData) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.t)
	stateSinkObject.Save(1, &s.tgstats)
	stateSinkObject.Save(2, &s.pidns)
}

func (s *taskStatData) afterLoad() {}

// +checklocksignore
func (s *taskStatData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.t)
	stateSourceObject.Load(1, &s.tgstats)
	stateSourceObject.Load(2, &s.pidns)
}

func (s *statmData) StateTypeName() string {
	return "pkg/sentry/fs/proc.statmData"
}

func (s *statmData) StateFields() []string {
	return []string{
		"t",
	}
}

func (s *statmData) beforeSave() {}

// +checklocksignore
func (s *statmData) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.t)
}

func (s *statmData) afterLoad() {}

// +checklocksignore
func (s *statmData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.t)
}

func (s *statusData) StateTypeName() string {
	return "pkg/sentry/fs/proc.statusData"
}

func (s *statusData) StateFields() []string {
	return []string{
		"t",
		"pidns",
	}
}

func (s *statusData) beforeSave() {}

// +checklocksignore
func (s *statusData) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.t)
	stateSinkObject.Save(1, &s.pidns)
}

func (s *statusData) afterLoad() {}

// +checklocksignore
func (s *statusData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.t)
	stateSourceObject.Load(1, &s.pidns)
}

func (i *ioData) StateTypeName() string {
	return "pkg/sentry/fs/proc.ioData"
}

func (i *ioData) StateFields() []string {
	return []string{
		"ioUsage",
	}
}

func (i *ioData) beforeSave() {}

// +checklocksignore
func (i *ioData) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.ioUsage)
}

func (i *ioData) afterLoad() {}

// +checklocksignore
func (i *ioData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.ioUsage)
}

func (c *comm) StateTypeName() string {
	return "pkg/sentry/fs/proc.comm"
}

func (c *comm) StateFields() []string {
	return []string{
		"SimpleFileInode",
		"t",
	}
}

func (c *comm) beforeSave() {}

// +checklocksignore
func (c *comm) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.SimpleFileInode)
	stateSinkObject.Save(1, &c.t)
}

func (c *comm) afterLoad() {}

// +checklocksignore
func (c *comm) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.SimpleFileInode)
	stateSourceObject.Load(1, &c.t)
}

func (f *commFile) StateTypeName() string {
	return "pkg/sentry/fs/proc.commFile"
}

func (f *commFile) StateFields() []string {
	return []string{
		"t",
	}
}

func (f *commFile) beforeSave() {}

// +checklocksignore
func (f *commFile) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.t)
}

func (f *commFile) afterLoad() {}

// +checklocksignore
func (f *commFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.t)
}

func (a *auxvec) StateTypeName() string {
	return "pkg/sentry/fs/proc.auxvec"
}

func (a *auxvec) StateFields() []string {
	return []string{
		"SimpleFileInode",
		"t",
	}
}

func (a *auxvec) beforeSave() {}

// +checklocksignore
func (a *auxvec) StateSave(stateSinkObject state.Sink) {
	a.beforeSave()
	stateSinkObject.Save(0, &a.SimpleFileInode)
	stateSinkObject.Save(1, &a.t)
}

func (a *auxvec) afterLoad() {}

// +checklocksignore
func (a *auxvec) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &a.SimpleFileInode)
	stateSourceObject.Load(1, &a.t)
}

func (f *auxvecFile) StateTypeName() string {
	return "pkg/sentry/fs/proc.auxvecFile"
}

func (f *auxvecFile) StateFields() []string {
	return []string{
		"t",
	}
}

func (f *auxvecFile) beforeSave() {}

// +checklocksignore
func (f *auxvecFile) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.t)
}

func (f *auxvecFile) afterLoad() {}

// +checklocksignore
func (f *auxvecFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.t)
}

func (o *oomScoreAdj) StateTypeName() string {
	return "pkg/sentry/fs/proc.oomScoreAdj"
}

func (o *oomScoreAdj) StateFields() []string {
	return []string{
		"SimpleFileInode",
		"t",
	}
}

func (o *oomScoreAdj) beforeSave() {}

// +checklocksignore
func (o *oomScoreAdj) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
	stateSinkObject.Save(0, &o.SimpleFileInode)
	stateSinkObject.Save(1, &o.t)
}

func (o *oomScoreAdj) afterLoad() {}

// +checklocksignore
func (o *oomScoreAdj) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &o.SimpleFileInode)
	stateSourceObject.Load(1, &o.t)
}

func (f *oomScoreAdjFile) StateTypeName() string {
	return "pkg/sentry/fs/proc.oomScoreAdjFile"
}

func (f *oomScoreAdjFile) StateFields() []string {
	return []string{
		"t",
	}
}

func (f *oomScoreAdjFile) beforeSave() {}

// +checklocksignore
func (f *oomScoreAdjFile) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.t)
}

func (f *oomScoreAdjFile) afterLoad() {}

// +checklocksignore
func (f *oomScoreAdjFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.t)
}

func (imio *idMapInodeOperations) StateTypeName() string {
	return "pkg/sentry/fs/proc.idMapInodeOperations"
}

func (imio *idMapInodeOperations) StateFields() []string {
	return []string{
		"InodeSimpleAttributes",
		"InodeSimpleExtendedAttributes",
		"t",
		"gids",
	}
}

func (imio *idMapInodeOperations) beforeSave() {}

// +checklocksignore
func (imio *idMapInodeOperations) StateSave(stateSinkObject state.Sink) {
	imio.beforeSave()
	stateSinkObject.Save(0, &imio.InodeSimpleAttributes)
	stateSinkObject.Save(1, &imio.InodeSimpleExtendedAttributes)
	stateSinkObject.Save(2, &imio.t)
	stateSinkObject.Save(3, &imio.gids)
}

func (imio *idMapInodeOperations) afterLoad() {}

// +checklocksignore
func (imio *idMapInodeOperations) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &imio.InodeSimpleAttributes)
	stateSourceObject.Load(1, &imio.InodeSimpleExtendedAttributes)
	stateSourceObject.Load(2, &imio.t)
	stateSourceObject.Load(3, &imio.gids)
}

func (imfo *idMapFileOperations) StateTypeName() string {
	return "pkg/sentry/fs/proc.idMapFileOperations"
}

func (imfo *idMapFileOperations) StateFields() []string {
	return []string{
		"iops",
	}
}

func (imfo *idMapFileOperations) beforeSave() {}

// +checklocksignore
func (imfo *idMapFileOperations) StateSave(stateSinkObject state.Sink) {
	imfo.beforeSave()
	stateSinkObject.Save(0, &imfo.iops)
}

func (imfo *idMapFileOperations) afterLoad() {}

// +checklocksignore
func (imfo *idMapFileOperations) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &imfo.iops)
}

func (u *uptime) StateTypeName() string {
	return "pkg/sentry/fs/proc.uptime"
}

func (u *uptime) StateFields() []string {
	return []string{
		"SimpleFileInode",
		"startTime",
	}
}

func (u *uptime) beforeSave() {}

// +checklocksignore
func (u *uptime) StateSave(stateSinkObject state.Sink) {
	u.beforeSave()
	stateSinkObject.Save(0, &u.SimpleFileInode)
	stateSinkObject.Save(1, &u.startTime)
}

func (u *uptime) afterLoad() {}

// +checklocksignore
func (u *uptime) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &u.SimpleFileInode)
	stateSourceObject.Load(1, &u.startTime)
}

func (f *uptimeFile) StateTypeName() string {
	return "pkg/sentry/fs/proc.uptimeFile"
}

func (f *uptimeFile) StateFields() []string {
	return []string{
		"startTime",
	}
}

func (f *uptimeFile) beforeSave() {}

// +checklocksignore
func (f *uptimeFile) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.startTime)
}

func (f *uptimeFile) afterLoad() {}

// +checklocksignore
func (f *uptimeFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.startTime)
}

func (v *versionData) StateTypeName() string {
	return "pkg/sentry/fs/proc.versionData"
}

func (v *versionData) StateFields() []string {
	return []string{
		"k",
	}
}

func (v *versionData) beforeSave() {}

// +checklocksignore
func (v *versionData) StateSave(stateSinkObject state.Sink) {
	v.beforeSave()
	stateSinkObject.Save(0, &v.k)
}

func (v *versionData) afterLoad() {}

// +checklocksignore
func (v *versionData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &v.k)
}

func init() {
	state.Register((*execArgInode)(nil))
	state.Register((*execArgFile)(nil))
	state.Register((*fdDir)(nil))
	state.Register((*fdDirFile)(nil))
	state.Register((*fdInfoDir)(nil))
	state.Register((*filesystemsData)(nil))
	state.Register((*filesystem)(nil))
	state.Register((*taskOwnedInodeOps)(nil))
	state.Register((*staticFileInodeOps)(nil))
	state.Register((*loadavgData)(nil))
	state.Register((*meminfoData)(nil))
	state.Register((*mountInfoFile)(nil))
	state.Register((*mountsFile)(nil))
	state.Register((*ifinet6)(nil))
	state.Register((*netDev)(nil))
	state.Register((*netSnmp)(nil))
	state.Register((*netRoute)(nil))
	state.Register((*netUnix)(nil))
	state.Register((*netTCP)(nil))
	state.Register((*netTCP6)(nil))
	state.Register((*netUDP)(nil))
	state.Register((*proc)(nil))
	state.Register((*self)(nil))
	state.Register((*threadSelf)(nil))
	state.Register((*rootProcFile)(nil))
	state.Register((*statData)(nil))
	state.Register((*mmapMinAddrData)(nil))
	state.Register((*overcommitMemory)(nil))
	state.Register((*maxMapCount)(nil))
	state.Register((*hostname)(nil))
	state.Register((*hostnameFile)(nil))
	state.Register((*tcpMemInode)(nil))
	state.Register((*tcpMemFile)(nil))
	state.Register((*tcpSack)(nil))
	state.Register((*tcpSackFile)(nil))
	state.Register((*tcpRecovery)(nil))
	state.Register((*tcpRecoveryFile)(nil))
	state.Register((*ipForwarding)(nil))
	state.Register((*ipForwardingFile)(nil))
	state.Register((*portRangeInode)(nil))
	state.Register((*portRangeFile)(nil))
	state.Register((*taskDir)(nil))
	state.Register((*subtasks)(nil))
	state.Register((*subtasksFile)(nil))
	state.Register((*exe)(nil))
	state.Register((*cwd)(nil))
	state.Register((*namespaceSymlink)(nil))
	state.Register((*memData)(nil))
	state.Register((*memDataFile)(nil))
	state.Register((*mapsData)(nil))
	state.Register((*smapsData)(nil))
	state.Register((*taskStatData)(nil))
	state.Register((*statmData)(nil))
	state.Register((*statusData)(nil))
	state.Register((*ioData)(nil))
	state.Register((*comm)(nil))
	state.Register((*commFile)(nil))
	state.Register((*auxvec)(nil))
	state.Register((*auxvecFile)(nil))
	state.Register((*oomScoreAdj)(nil))
	state.Register((*oomScoreAdjFile)(nil))
	state.Register((*idMapInodeOperations)(nil))
	state.Register((*idMapFileOperations)(nil))
	state.Register((*uptime)(nil))
	state.Register((*uptimeFile)(nil))
	state.Register((*versionData)(nil))
}
